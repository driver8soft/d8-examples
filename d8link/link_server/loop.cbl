       IDENTIFICATION DIVISION.
       PROGRAM-ID. loop.
       ENVIRONMENT DIVISION.
       CONFIGURATION SECTION.
       DATA DIVISION.
       FILE SECTION.
       WORKING-STORAGE SECTION.
      * Declare variables in the WORKING-STORAGE section
       01 TOTAL PIC S9(9) VALUE ZEROES.
       LINKAGE SECTION.
      * Data to share with COBOL subroutines 
       01 COMMAREA.
          05 N-TIMES PIC S9(9) COMP.
      * Data to share with COBOL subroutines 
       PROCEDURE DIVISION USING COMMAREA. 
      * code goes here!
           DISPLAY "Starting loop".
           PERFORM PROCESS N-TIMES TIMES.
           DISPLAY "Total: " TOTAL.
           GOBACK.

           PROCESS.
           CALL "routine" USING TOTAL.

