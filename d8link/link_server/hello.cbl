       IDENTIFICATION DIVISION.
       PROGRAM-ID. hello.
       ENVIRONMENT DIVISION.
       CONFIGURATION SECTION.
       DATA DIVISION.
       FILE SECTION.
       WORKING-STORAGE SECTION.
      * In COBOL, you declare variables in the WORKING-STORAGE section

       LINKAGE SECTION.
      * Data to share with COBOL subroutines 
       01 COMMAREA PIC X(20).
       PROCEDURE DIVISION USING COMMAREA. 
      * code goes here!
           DISPLAY "Hello, " COMMAREA.
           GOBACK.


