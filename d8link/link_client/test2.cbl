       IDENTIFICATION DIVISION.
       PROGRAM-ID. test2.
       ENVIRONMENT DIVISION.
       CONFIGURATION SECTION.
       DATA DIVISION.
       WORKING-STORAGE SECTION.
      * In COBOL, you declare variables in the WORKING-STORAGE section
       01 PROG-NAME PIC X(8) VALUE "say".
       01 COMMLEN PIC 9(9) COMP.
       01 COMMAREA.
           05 A PIC X(13).
           05 B PIC S9(7)V9(2).
           05 C PIC S9(7)V9(2) COMP-3.
           05 D PIC S9(4) COMP.
           05 E PIC S9(9) COMP.
       LINKAGE SECTION.
      * Data to share with COBOL subroutines 
      
       PROCEDURE DIVISION.
           
           INITIALIZE COMMAREA.

           COMPUTE COMMLEN = LENGTH OF COMMAREA.
           CALL "D8link" USING PROG-NAME COMMAREA COMMLEN.
           
           DISPLAY "COBOL A: " A.
           DISPLAY "COBOL B: " B.
           DISPLAY "COBOL C: " C.
           DISPLAY "COBOL D: " D.
           DISPLAY "COBOL E: " E.

           DISPLAY "COBOL RETURN-CODE: " RETURN-CODE.

           GOBACK.
           